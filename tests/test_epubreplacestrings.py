#!/usr/bin/env python

from importlib.machinery import ModuleSpec, SourceFileLoader
from importlib.util import spec_from_loader, module_from_spec
import os.path
import types
import unittest
from dataclasses import dataclass


def import_from_source(name: str, file_path: str) -> types.ModuleType:
    loader: SourceFileLoader = SourceFileLoader(name, file_path)
    spec: ModuleSpec = spec_from_loader(loader.name, loader)
    module: types.ModuleType = module_from_spec(spec)
    loader.exec_module(module)
    return module


script_path: str = os.path.abspath(
    os.path.join(
        os.path.dirname(os.path.abspath(__file__)),
        "..",
        "bin",
        "epubreplacestrings",
    )
)

epub_replace_string: types.ModuleType = import_from_source("epubreplacestring", script_path)


class EpubStringReplaceTestCase(unittest.TestCase):
    # testRegexReplace is based on https://lorenzopeppoloni.com/tabledriventestspy/
    def testRegexReplace(self):
        @dataclass
        class TestCase:
            name: str
            input: str
            expected: str

        testcases = [
            TestCase(
                name="make sure that html comments are left alone",
                input="<!--this is a comment. comments are not displayed in the browser-->",
                expected="<!--this is a comment. comments are not displayed in the browser-->",
            ),
            TestCase(
                name="make sure that two en dashes are replaced with an em dash",
                input="-- test --",
                expected="— test —",
            ),
            TestCase(
                name="make sure that html comments are left alone",
                input="""
                  ...
                  . . .
                  . .. 
                  .. .
                  .  . .
                """,
                expected="""
                  …
                  …
                  … 
                  …
                  .  . .
                """,
            ),
            TestCase(
                name="make sure that a lowercase 'by the by' results in a lowercase 'by the way'",
                input="by the by",
                expected="by the way",
            ),
            TestCase(
                name="make sure that an uppercase 'By the by' results in an uppercase 'By the way'",
                input="By the by",
                expected="By the way",
            ),
            TestCase(
                name="make sure that an uppercase 'Sneaked' results in an uppercase 'Snuck'",
                input="Sneaked",
                expected="Snuck",
            ),
            TestCase(
                name="make sure that a lowercase 'snuck' results in a lowercase 'snuck'",
                input="On his way he sneaked out the door",
                expected="On his way he snuck out the door",
            ),
            TestCase(
                name="make sure that single tilde is converted to an exclamation mark",
                input="~wow isn't this a joy~",
                expected="!wow isn't this a joy!",
            ),
            TestCase(
                name="make sure that multiple tildes in a row are not converted to an exclamation mark",
                input="~~ is completely ~~~ left alone",
                expected="~~ is completely ~~~ left alone",
            ),
            TestCase(
                name="make sure that a lowercase 'a bolt out of the blue' is correctly converted to 'out of the blue",
                input="a bolt out of the blue",
                expected="out of the blue",
            ),
            TestCase(
                name="make sure that an uppercase 'A bolt out of the blue' is correctly converted to 'Out of the blue",
                input="A bolt out of the blue, a loud bang went off waking up people who had been sleeping soundly up until that point.",
                expected="Out of the blue, a loud bang went off waking up people who had been sleeping soundly up until that point.",
            ),
            TestCase(
                name="make sure that a lowercase 'little wonder' is correctly converted to 'no wonder",
                input="little wonder your attempt failed",
                expected="no wonder your attempt failed",
            ),
            TestCase(
                name="make sure that an uppercase 'Little wonder' is correctly converted to 'No wonder",
                input="Little wonder, you were outmatched from the start",
                expected="No wonder, you were outmatched from the start",
            ),
        ]

        for case in testcases:
            actual = epub_replace_string.replace_strings(case.input)
            self.assertEqual(
                case.expected,
                actual,
                "failed test {} expected {}, actual {}".format(
                    case.name, case.expected, actual
                ),
            )

    def testReplacementParser(self):
        @dataclass
        class TestCase:
            name: str
            input: str
            expected: dict[str, int]

        testcases = [
            TestCase(
                name="make sure that an empty table results in an empty dictionary",
                input="""| Text to replace | Text replacement |
                | ---- | ---- |
                """,
                expected={},
            ),
            TestCase(
                name="make sure that a non-empty table results in the appropriate amount of entries being placed in a dictionary",
                input="""| Text to replace | Text replacement |
                | ---- | ---- |
                | replace | with me |
                | "I am quoted" | 'I am single quoted' |
                """,
                expected={
                    'replace': 'with me',
                    '\"I am quoted\"': '\'I am single quoted\'',
                },
            ),
        ]

        for case in testcases:
            actual = epub_replace_string.parse_text_replacements(case.input)
            self.assertEqual(
                case.expected,
                actual,
                "failed test {} expected {}, actual {}".format(
                    case.name, case.expected, actual
                ),
            )
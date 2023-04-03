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

replace_strings: types.ModuleType = import_from_source("epubreplacestring", script_path)


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
        ]

        for case in testcases:
            actual = replace_strings.replace_strings(case.input)
            self.assertEqual(
                case.expected,
                actual,
                "failed test {} expected {}, actual {}".format(
                    case.name, case.expected, actual
                ),
            )

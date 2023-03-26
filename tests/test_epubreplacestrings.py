#!/usr/bin/env python

from importlib.machinery import ModuleSpec, SourceFileLoader
from importlib.util import spec_from_loader, module_from_spec
import os.path
import types
import unittest


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
    def testHtmlCommentLeftAlone(self):
        input = "<!--This is a comment. Comments are not displayed in the browser-->"
        output = replace_strings.replace_strings(input)

        assert output == input

    def testTwoEnDashesReplacedWithEmDash(self):
        input = "-- test --"
        output = replace_strings.replace_strings(input)

        assert output == "— test —"

    def testMulitplePeriodsConvertedToProperEllipsis(self):
        input = """
        ...
        . . .
        . .. 
        .. .
        .  . .
        """
        expected_output = """
        …
        …
        … 
        …
        .  . .
        """
        output = replace_strings.replace_strings(input)

        assert output == expected_output

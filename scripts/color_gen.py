#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import argparse
import csv
import sys
import typing

__author__ = "Benjamin Kane"
__version__ = "0.1.0"
__doc__ = f"""
<description>
Examples:
    {sys.argv[0]}
Help:
Please see Benjamin Kane for help.
Code at <repo>
"""


def parse_args(*args, **kwargs):
    parser = argparse.ArgumentParser(
        description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter
    )

    # Use a file or stdin for an argument
    # https://stackoverflow.com/a/11038508/2958070
    parser.add_argument(
        "infile",
        nargs="?",
        type=argparse.FileType("r"),
        default=sys.stdin,
        help="data from https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences?redirectedfrom=MSDN#text-formatting , copied and pasted into a TSV",
    )
    parser.add_argument(
        "--max_rows",
        type=int,
        default=1000,
        help="max number of TSV rows to consider. Useful for testing",
    )

    return parser.parse_args(*args, **kwargs)


class ColorInfo(typing.NamedTuple):
    code: str
    name: str
    description: str


def main():
    replacements = {
        "Bold/Bright": "Bold",
        "Nobold/bright": "Nobold",
        "Positive(Nonegative)": "Positive",
    }
    args = parse_args()
    with args.infile:
        reader = csv.reader(args.infile, delimiter="\t")
        colorinfos = []
        for row in reader:
            code = row[0].strip()
            name = row[1].strip().replace(" ", "")
            description = row[2].strip()
            if name in replacements:
                name = replacements[name]

            colorinfos.append(ColorInfo(code=code, name=name, description=description))

    print("package color")
    print()

    print("// These will be set to the proper color codes if Enable() returns successfully. ")
    print("// generated from gen_colors")
    print("var (")
    for i, ci in enumerate(colorinfos):
        print(f"\t// {ci.name} - {ci.description}")
        print(f'\t{ci.name} = ""')
        print()

        if i > args.max_rows:
            break
    print(")")
    print()

    print("func setColors() {")
    for i, ci in enumerate(colorinfos):
        print(f'\t{ci.name} = "\\033[{ci.code}m"')

        if i > args.max_rows:
            break

    print("}")


if __name__ == "__main__":
    main()

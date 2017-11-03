package python

var setupyText = `import re
from os import path
from codecs import open  # To use a consistent encoding
from setuptools import setup, find_packages

here = path.abspath(path.dirname(__file__))

# Get the long description from the relevant file
with open(path.join(here, 'README.md'), encoding='utf-8') as f:
    long_description = f.read()
# Get version without importing, which avoids dependency issues


def get_version():
    with open('%s/__init__.py') as version_file:
        return re.search(r"""__version__\s+=\s+(['"])(?P<version>.+?)\1""",
                         version_file.read()).group('version')

install_requires = ['future']
test_requires = ['pytest', 'pytest-sugar', 'pytest-asyncio', 'pytest-cov', ]
setup(
    name='%s',
    description="Some description about your project",
    long_description=long_description,
    version=get_version(),
    include_package_data=True,
    install_requires=install_requires,
    setup_requires=['pytest-runner'],
    tests_require=test_requires,
    packages=find_packages(),
    zip_safe=False,
    author="Your name here",
    download_url="your project url/archive/{}.tar.gz".format(get_version()),
    classifiers=[
        "Programming Language :: Python :: 2",
        "Programming Language :: Python :: 2.7",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.4",
        "Programming Language :: Python :: 3.5",
        "Programming Language :: Python :: 3.6", ]
)`

var setupCfgText = `[aliases]
test=pytest

[tool:pytest]
addopts = --verbose -vv --cov %s`

var conftestText = `# !/usr/bin/env python
# -*- coding: utf-8 -*-
"""Bunch of fixtures to be used across the tests."""

import pytest

@pytest.fixture(scope="function")
def test_fixture(request):
    """Create a test fixture."""

    myvar = 5

    def tear_down():
        # clean up here
        pass
    request.addfinalizer(tear_down)

    return myvar`

var testfileText = `# !/usr/bin/env python
# -*- coding: utf-8 -*-
"""Test modules."""


def test_list_files_older_than(test_fixture):
    """Run a test."""

    assert(5 == test_fixture)`

var initpyText = `__version__ = "0.1.0"`

var makefileText = `.DEFAULT_GOAL := help

help:          ## Show available options with this Makefile
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY : test
test:          ## Run all the tests
test:
	python setup.py test`

var requirementsText = ``

var readmeText = `%s:
---

This project can be used to ...

Install:
---
1. Clone the project.

2. Create a virtualenv using venv or virtualenv:` +
	"\n\n\t```\n" +
	"\tcd %s\n" +
	"\t. venv/bin/activate\n" +
	"\t```\n" +

	`3. If using conda, edit dev-environment.yml file:` +

	"\n\n\t```\n" +
	"\tconda env create --force -f dev_environment.yml\n" +
	"\tsource activate %s\n" +
	"\t```\n\n" +

	`Test:
---
1. To run the tests:` +
	"\n\n\t`make test`\n\n" +

	`
Usage:
---

....

Example:
---

...
`

var manifestText = `
include README.md`

var devEnvYamlText = `name: %s
channels:
  - defaults
dependencies:
- python=3.6.3
- pip:
    - future`

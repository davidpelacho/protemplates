package python

const travisText = `language: python
python:
  - "2.7"
  - "3.5"
  - "3.6"
# command to install dependencies
install:
  - pip install -r requirements.txt
# command to run tests
script:
  - python setup.py test`

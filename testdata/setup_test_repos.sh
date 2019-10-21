#!/bin/bash

rm -rf testdata/commits_on_branch
rm -rf testdata/git_tags
rm -rf testdata/detached_tags

  cd testdata

  echo 'testdata/commits_on_branch_test/ directory does not exist at the root; creating...'
  git clone commits_on_branch.bundle
  echo 'done'

  echo 'testdata/git_tags/ directory does not exist at the root; creating...'
  git clone git_tags.bundle
  echo 'done'


echo 'testdata/detached_tags/ directory does not exist at the root; creating...'
  git clone detached_tags.bundle
  echo 'done'

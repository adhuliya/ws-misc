# a set of utility functions

from typing import Dict, Tuple
from .person import Person

def backupdb(people):
    # TODO
    pass

def create_passwd_dict(people : Dict[int, Person]) -> Dict[str, Tuple[int, str]]:
  d = {}
  for k in people:
    person = people[k]
    d[person.email_id] = (k, person.passwd)

  return d



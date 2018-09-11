from .person import Person
from typing import List

class PageData:
    """
    Object of this class is sent to page templates
    for rendering.
    """
    def __init__(self,
            person : Person = None,
            relatives : List[Person] = None
            ):
        self.person = person
        self.relatives = relatives
        self.profile_pic = None
        self.avatar_pic = None
        self.pics = None
        self.family = None



#!/usr/bin/env python3

# A python class that defines a person and its associated data

from typing import Tuple, List, Dict, Set

class Person:
    def __init__(self,
            person_id : int,
            real_entry:bool = False,
            name : str = None,
            nick_name : str = None,
            gender : str = 'M', # M or F
            dob : Tuple[int,int,int] = None, # year month day
            profile_pic: str = None, pics: List[str] = None,
            mother_id : int = None,
            father_id : int = None,
            spouse_id : int = None,
            friends_id : Set[int] = None,
            children_id : Set[int] = None,
            email_id  : str = None,  # also the userid
            passwd : str = None,
            twitter_id : str = None,
            insta_id : str = None,
            fb_id : str = None,
            profession : str = None,  # one line profession description
            aboutme     : str = None, # markdown file containing info
        ) -> None:

        self.person_id      = person_id
        self.real_entry     = real_entry
        self.name           = name
        self.nick_name      = nick_name
        self.gender         = gender
        self.dob            = dob
        self.profile_pic    = profile_pic
        self.pics           = pics
        self.mother_id      = mother_id 
        self.father_id      = father_id
        self.spouse_id      = spouse_id
        self.children_id    = children_id
        self.friends_id     = friends_id
        self.email_id       = email_id
        self.passwd         = passwd
        self.twitter_id     = twitter_id
        self.insta_id       = insta_id
        self.fb_id          = fb_id
        self.profession     = profession
        self.aboutme        = aboutme


    def __str__(self):
        return self.__repr__()

    def __repr__(self):
        return f"Person (person_id = {self.person_id}, name = '{self.name}', nick_name = '{self.nick_name}', gender = '{self.gender}', dob = {self.dob}, profile_pic = '{self.profile_pic}', pics = {self.pics}, mother_id = {self.mother_id}, father_id = {self.father_id}, spouse_id = {self.spouse_id}, children_id = {self.children_id}, friends_id = {self.friends_id}, email_id = '{self.email_id}', passwd = '{self.passwd}', twitter_id = '{self.twitter_id}', fb_id = '{self.fb_id}', insta_id = '{self.insta_id}', real_entry = {self.real_entry}, profession = '{self.profession}', aboutme = '{self.aboutme}')" 


if __name__ == "__main__":
    print( Person(
        real_entry = True,
        person_id = 1,
        nick_name = "anshu",
        name = "anshuman dhuliya",
        gender = "M",
        dob = (0, 8, 18),
        profile_pic = "000001-0001.jpg", # format: <person_id>-<pic_id>.jpg
        pics = ["000001-0001.jpg", "000001-0001.jpg"],
        mother_id = None,
        father_id = None,
        spouse_id = None,
        friends_id = None,
        email_id = "anshumandhuliya@gmail.com",
        passwd = "anshuisneo",
        aboutme = "dig in to dig out",
        twitter_id = None,
        fb_id = None,
        insta_id = None
        ))


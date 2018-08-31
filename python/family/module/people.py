#!/usr/bin/env python3

# list of persons, relation between them defines a family
from typing import Dict
from .person import Person

people: Dict[int, Person] = {
        1:
    Person(
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
        children_id = None,
        friends_id = None,
        email_id = "anshumandhuliya@gmail.com",
        passwd = "anshuisneo",
        profession = "Research Scholar",
        aboutme = "000001.md",
        twitter_id = None,
        fb_id = None,
        insta_id = None
        ),
}


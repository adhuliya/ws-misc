#!/usr/bin/env python3

# list of persons, relation between them defines a family
from typing import Dict
from .person import Person

people: Dict[int, Person] = {

        1:
    Person(
        person_id = 1,  # should match dict key
        real_entry = True,
        name = "priyanka dhuliya",
        nick_name = "rimi",
        gender = "M",
        dob = (0, 6, 2),
        avatar_pic = None, # e.g. "000001-avatar.jpg",
        profile_pic = None, # e.g. "000001-profile.jpg",
        # pic name format: <person_id>-<pic_id>.jpg : e.g. ["000001-001.jpg", ...]
        pics = None, # list of pics
        mother_id = 3,
        father_id = 4,
        spouse_id = 2,
        children_id = None,
        friends_id = None,
        email_id = "iampriyanka.dhuliya@gmail.com",
        passwd = "family",
        twitter_id = None,
        fb_id = None,
        insta_id = None,
        profession = "Home Manager",
        motd = None,  # e.g. "Life is Beautiful!",
        about_me = "000001.md",
        ),

        2:
    Person(
        person_id = 2,  # should match dict key
        real_entry = True,
        name = "anshuman dhuliya",
        nick_name = "anshu",
        gender = "M",
        dob = (0, 8, 18),
        avatar_pic = None, # e.g. "000001-avatar.jpg",
        profile_pic = None, # e.g. "000001-profile.jpg",
        # pic name format: <person_id>-<pic_id>.jpg : e.g. ["000001-001.jpg", ...]
        pics = None, # list of pics
        mother_id = None,
        father_id = None,
        spouse_id = 1,
        children_id = None,
        friends_id = None,
        email_id = "anshumandhuliya@gmail.com",
        passwd = "anshuisneo",
        twitter_id = None,
        fb_id = None,
        insta_id = None,
        profession = "Research Scholar, IIT Bombay",
        motd = None,  # e.g. "Life is Beautiful!",
        about_me = "000001.md",
        ),

        3:
    Person(
        person_id = 3,  # should match dict key
        real_entry = True,
        name = "monika banerjee",
        nick_name = "mithali",
        gender = "F",
        dob = (0, 6, 15),
        avatar_pic = None, # e.g. "000001-avatar.jpg",
        profile_pic = None, # e.g. "000001-profile.jpg",
        # pic name format: <person_id>-<pic_id>.jpg : e.g. ["000001-001.jpg", ...]
        pics = None, # list of pics
        mother_id = None,
        father_id = None,
        spouse_id = 4,
        children_id = [1],
        friends_id = None,
        email_id = "iammonika.banerjee@gmail.com",
        passwd = "family",
        twitter_id = None,
        fb_id = None,
        insta_id = None,
        profession = "Home Manager",
        motd = None,  # e.g. "Life is Beautiful!",
        about_me = "000001.md",
        ),

        4:
    Person(
        person_id = 4,  # should match dict key
        real_entry = True,
        name = "chandan banerjee",
        nick_name = "chandan",
        gender = "M",
        dob = (0, 12, 17),
        avatar_pic = None, # e.g. "000001-avatar.jpg",
        profile_pic = None, # e.g. "000001-profile.jpg",
        # pic name format: <person_id>-<pic_id>.jpg : e.g. ["000001-001.jpg", ...]
        pics = None, # list of pics
        mother_id = None,
        father_id = None,
        spouse_id = 3,
        children_id = [1],
        friends_id = None,
        email_id = "banerjeegrandpiano@gmail.com",
        passwd = "family",
        twitter_id = None,
        fb_id = None,
        insta_id = None,
        profession = "Accountant, DLW Varanasi",
        motd = None,  # e.g. "Life is Beautiful!",
        about_me = "000001.md",
        ),

        5:
    Person(
        person_id = 5,  # should match dict key
        real_entry = True,
        name = "anubha dhulia",
        nick_name = "anu",
        gender = "F",
        dob = (0, 11, 18),
        avatar_pic = None, # e.g. "000001-avatar.jpg",
        profile_pic = None, # e.g. "000001-profile.jpg",
        # pic name format: <person_id>-<pic_id>.jpg : e.g. ["000001-001.jpg", ...]
        pics = None, # list of pics
        mother_id = None,
        father_id = None,
        spouse_id = None,
        children_id = [],
        friends_id = None,
        email_id = "anubha.dhuliya@gmail.com",
        passwd = "family",
        twitter_id = None,
        fb_id = None,
        insta_id = None,
        profession = "Advocate, International Arbitration",
        motd = None,  # e.g. "Life is Beautiful!",
        about_me = "000001.md",
        ),
}


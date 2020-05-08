import datetime

from django.db import models
from django.utils import timezone

# Create your models here.
# Useful command sequence when model is created or changed:
#   py3 manage.py makemigrations polls    # generate the migrations
#   py3 manage.py sqlmigrate polls 0001   # see the sql generated (optional)
#   py3 manage.py check                   # check for inconsistencies (optional)
#   py3 manage.py migrate                 # migrates the database

class Question(models.Model):
    question_text = models.CharField(max_length=200)
    pub_date = models.DateTimeField('date published')

    def __str__(self):
        return self.question_text

    def was_published_recently(self):
        now = timezone.now()
        return now >= self.pub_date >= now - datetime.timedelta(days=1)


class Choice(models.Model):
    question = models.ForeignKey(Question, on_delete=models.CASCADE)
    choice_text = models.CharField(max_length=200)
    votes = models.IntegerField(default=0)

    def __str__(self):
        return self.choice_text


class User(models.Model):
    name       = models.CharField('Real User Name', max_length=64)
    name_pub   = models.CharField('Public User Name', max_length=16)
    email      = models.CharField(max_length=50)
    isd_code   = models.CharField(max_length=4)
    mobile     = models.CharField(max_length=10)
    country    = models.CharField(max_length=20)
    active     = models.IntegerField(default=1)

    
class Event(models.Model):
    name      = models.CharField(max_length=128)
    name_pub  = models.CharField(max_length=16)
    startsAt  = models.DateTimeField('Event Start')
    endsAt    = models.DateTimeField('Event End')
    creator   = models.ForeignKey(User, on_delete=models.CASCADE)
    active    = models.IntegerField(default=0)
    iteration = models.IntegerField(default=0)


class Token(models.Model):
    user    = models.ForeignKey(User, on_delete=models.CASCADE)
    event   = models.ForeignKey(Event, on_delete=models.CASCADE)
    iteration = models.IntegerField(default=0) # event iteration the token belongs to
    toknum = models.IntegerField('Token Number', default=0)
    active = models.IntegerField(default=1)
    created  = models.DateTimeField('Token Issue Time')
    message = models.CharField(max_length=128)





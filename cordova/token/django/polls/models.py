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



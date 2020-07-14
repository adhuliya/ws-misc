import datetime

from django.db    import models
from django.utils import timezone

#AD: Some Notes
# 1. Change your models (in models.py).
# 2. Run `python manage.py makemigrations [polls]`
#    to create migrations for those changes
# 3. Run `python manage.py migrate` to apply those changes to the database.
# 4. View migration 0001: `python manage.py sqlmigrate polls 0001`
# 5. Check for any problems: `python manage.py check`


class Question(models.Model):
  question_text = models.CharField(max_length=200)
  pub_date      = models.DateTimeField('date published')


  def was_published_recently(self):
    return self.pub_date >= timezone.now() - datetime.timedelta(days=1)


  def __str__(self):
    return self.question_text


class Choice(models.Model):
  question    = models.ForeignKey(Question, on_delete=models.CASCADE)
  choice_text = models.CharField(max_length=200)
  votes       = models.IntegerField(default=0)


  def __str__(self):
    return self.choice_text



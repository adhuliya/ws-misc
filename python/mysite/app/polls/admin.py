from django.contrib import admin

#AD: First create the superuser:
# > $ py manage.py createsuperuser                                                                                            [±master ●●●]
# Username: admin
# Email address: lazynintel@gmail.com
# Password:
# Password (again):
# Superuser created successfully.

# Access Admin at: http://127.0.0.1:8000/admin/

# Register your models here.

from .models import Question, Choice

admin.site.register(Question)
admin.site.register(Choice)





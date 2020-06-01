from django.shortcuts import render, get_object_or_404
from django.http import HttpResponse, Http404, HttpResponseRedirect, JsonResponse
from django.template import loader
from django.urls import reverse
from django.db.models import F
from django.utils import timezone
import json

from .models import Question, Choice

# Create your views here.

def index(request):
  """Using render shortcut"""
  latest_question_list = Question.objects.order_by('-pub_date')[:5]
  context = {'latest_question_list': latest_question_list}
  return render(request, 'polls/index.html', context)

# def index(request):
#   """Without the render shortcut"""
#   latest_question_list = Question.objects.order_by('-pub_date')[:5]
#   template = loader.get_template('polls/index.html')
#   context = {
#     'latest_question_list': latest_question_list,
#   }
#   return HttpResponse(template.render(context, request))

# def index(request):
#   """Most basic view"""
#   return HttpResponse("Hello World! -- polls app index")

def detail(request, question_id):
  question = get_object_or_404(Question, pk=question_id)
  return render(request, 'polls/detail.html', {'question': question})

# def detail(request, question_id):
#   """without get_object_or_404() shortcut"""
#   try:
#     question = Question.objects.get(pk=question_id)
#   except Question.DoesNotExist:
#     raise Http404("Question does not exist")
#   return render(request, 'polls/detail.html', {'question': question})


def results(request, question_id):
  question = get_object_or_404(Question, pk=question_id)
  return render(request, 'polls/results.html', {'question': question})


def vote(request, question_id):
  question = get_object_or_404(Question, pk=question_id)
  try:
    selected_choice = question.choice_set.get(pk=request.POST['choice'])
  except (KeyError, Choice.DoesNotExist):
    # Redisplay the question voting form.
    return render(request, 'polls/detail.html', {
      'question': question,
      'error_message': "You didn't select a choice.",
    })
  else:
    # Ref: https://docs.djangoproject.com/en/3.0/ref/models/expressions/
    #selected_choice.votes += 1
    selected_choice.votes = F('votes') + 1  # thread safe way
    selected_choice.save()
    #To access the new value saved this way, the object must be reloaded:
    #  selected_choice.refresh_from_db()

    # Always return an HttpResponseRedirect after successfully dealing
    # with POST data. This prevents data from being posted twice if a
    # user hits the Back button.
    return HttpResponseRedirect(reverse('polls:results', args=(question.id,)))


def protectFile(request, filepath):
  """Serves static files that are protected,
  and uses nginx to server the files very efficiently.
  Ref: https://wellfire.co/learn/nginx-django-x-accel-redirects/
  """
  response = HttpResponse()
  response["Content-Disposition"] = "attachment; filename={0}".format(filepath)
  response['X-Accel-Redirect'] = "/protected/{0}".format(filepath)
  return response


def jsonSend(request):
  """Return json response."""
  data = {
    'time': str(timezone.now()),
    'is_active': True,
    'count': 28
  }
  return JsonResponse(data)


def jsonReceive(request):
  """Sample view to receive json data."""
  if request.is_ajax():
    if request.method == 'POST':
      jsonData = json.loads(request.body)
      try:
        data = jsonData['data']
      except KeyError:
        Http404("Malformed data!")
      return HttpResponse("Got json data")
  return HttpResponse("No json data.")



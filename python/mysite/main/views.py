from django.shortcuts import render

# Create your views here.

def index(request):
  """Using render shortcut"""
  context = {}
  return render(request, 'main/index.html', context)


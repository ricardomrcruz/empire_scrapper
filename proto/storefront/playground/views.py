from django.shortcuts import render
from django.http import HttpResponse

from item.models import Category, Item


def index(request):
    items = Item.objects.filter(is_sold=False)[0:6]
    categories = Item.objects.all()

    return render(request, "playground/index.html", {
        'categories': categories,
        'items': items, 
    })


def contact(request):
    return render(request, "playground/contact.html")

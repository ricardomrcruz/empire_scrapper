from playwright.sync_api import sync_playwright
from selectolax.parser import HTMLParser
from dataclasses import dataclass
from rich import print 


@dataclass
class Item:
    asin: str
    title: str
    price: str


def get_html(page, asin):
    url = f"https://www.amazon.fr/dp/{asin}"
    page.goto(url)
    html = HTMLParser(page.content())
    return html


def parse_html(html, asin):
    item= Item(
        asin:asin, 
        title: html.css_firs("span#productTitle").text(strip=True),
        price: html.css_first("span.a-price.aok-align-center.reinventPricePriceToPayMargin.priceToPay").text(strip=True),
    )
    print(html.css_first("title").text())
    print(asin)
    return Item
   


def run():
    asin = "B08H93ZRK9"
    pw = sync_playwright().start()
    browser = pw.chromium.launch()
    page = browser.new_page()
    html = get_html(page, asin)
    product = parse_html(html, asin)
    print(product)



def main():
    run()


if __name__ == "__main__":
    main()

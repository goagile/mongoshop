const host = "127.0.0.1"
const port = "8080"

var url = (path) => `http://${host}:${port}/api/v1/${path}`

var indexMain = () => {
    getBooks(url('books'))
}

$(document).ready(indexMain)

var getBooks = (url) => {
    console.log("getBooks", url)
    $.ajax({
        url: url,
        success: function(data, status) {
            console.log("getBooks", "data", data)
            // showBooks(data)
            showBooks(booksForTest)
        },
        error: function(data, status) {
            console.log("getProduct", url, status)
            showBooksEmpty()
        }
    })
}

var showBooks = (data) => {
    if (notFoundBookItems(data)) {
        console.log("showBooks", "data", data)
        showBooksEmpty()
        return
    }

    var books = data.items;

    $("ul#books-list").show()
    $("span.books-empty").hide()

    var $books = $("ul#books")
    $books.html("")

    for (var i = 0; i < books.length; i++) {
        var b = books[i]
        $books.append(bookItem(b))
    }
}

var notFoundBookItems = (data) => {
    return (
        undefined == data ||
        undefined == data.items || 
        !data.items.length
    )
}

var showBooksEmpty = () => {
    $("ul#books-list").hide()
    $("span.books-empty").show()
}

var bookItem = (b) => `
    <li class="book-item">
        ${bookItemCover(b)}
        <a class="book-item-price" href="${b.href}">${b.price}</a>
        <a class="book-item-discount" href="${b.href}">${b.discount}</a>
        <a class="book-item-title" href="${b.href}">${b.title}</a>
        ${bookItemAuthors(b)}
        ${bookItemPuboffice(b)}
        <input class="book-item-to-cart" type="button" value="В КОРЗИНУ"></input>
    </li>`

var bookItemCover = (b) => `
    <a class="book-item-cover" href="${b.href}">
        <img class="book-cover-img" 
            src=${b.cover.href}
            alt=${b.cover.title}
            title=${b.cover.title}>
        </img>
    </a>`

var bookItemAuthors = (b) => {
    var s = ``
    var authors = b.authors
    for(var i = 0; i < authors.length; i++) {
        var a = authors[i]
        s += `<a class="book-item-author" href="${a.href}">${a.name}</a>`
    }
    return s
}

var bookItemPuboffice = (b) => {
    var p = b.puboffice
    return `
        <a class="book-item-puboffice" href="${p.href}">
            ${p.name}
        </a>`
}

var booksForTest = {
    "items": [
        {
            href: ".",
            cover: {
                href: "./img/razrabotka-web-react.jpg",
                title: "Вайс, Хортон - Разработка веб-приложений в ReactJS"
            },
            price: 1800,
            discount: 15,
            title: "Разработка веб-приложений в ReactJS",
            authors: [
                {
                    name: "Вайс",
                    href: "."
                },
                {
                    name: "Хортон",
                    href: "."
                }
            ],
            puboffice: {
                name: "ДМК-Пресс",
                href: "."
            }
        },
        {
            href: ".",
            cover: {
                href: "./img/nepreriv-razvertivanye.jpg",
                title: "Хамбл, Фарли - Непрерывное развертывание ПО. Автоматизация процессов сборки, тестирования и внедрения новых версий"
            },
            price: 3440,
            discount: 20,
            title: "Непрерывное развертывание ПО. Автоматизация процессов сборки, тестирования и внедрения новых версий",
            authors: [
                {
                    name: "Хамбл",
                    href: "."
                },
                {
                    name: "Фарли",
                    href: "."
                }
            ],
            puboffice: {
                name: "Вильямс",
                href: "."
            }
        }
    ]
}

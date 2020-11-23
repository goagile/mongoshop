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
            showBooks({})
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
        <a class="book-item-title" href=".">${b.title}</a>
        <a class="book-item-author" href=".">${b.author}</a>
    </li>`

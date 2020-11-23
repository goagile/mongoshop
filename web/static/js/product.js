// product js

var productmain = () => {
    getProduct("/product?slug=wheel-barrow-9092");
};

var getProduct = (url) => {
    $.ajax({
        url: url,
        success: function(data, status) {
            var p = JSON.parse(data)
            console.log("getProduct", url, "data", p);
            showProduct(p);
        },
        error: function(data, status) {
            console.log("getProduct", url, status);
        }
    });
};

var showProduct = (p) => {
    $("h2#product-name").html(p.name);
    $("p#product-description").html(p.description);
    $("ul.product-details").append(detailsLine("weight", p.details.weight));
};

var detailsLine = (k, v) => 
    `<li>
        <span class="details-line-key">${k}</span>
        <span class="details-line-value">${v}</span>
    </li>`;

$(document).ready(productmain);

/* <li>
    <span class="details-line-key">weight</span>
    <span class="details-line-value">47</span>
</li>
<li>
    <span class="details-line-key">weight units</span>
    <span class="details-line-value">lbs</span>
</li>
<li>
    <span class="details-line-key">model num</span>
    <span class="details-line-value">lbs</span>
</li>
<li>
    <span class="details-line-key">manufacturer</span>
    <span class="details-line-value">Acme</span>
</li>
<li>
    <span class="details-line-key">color</span>
    <span class="details-line-value">Green</span>
</li> */
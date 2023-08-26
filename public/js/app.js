function setSearchProvider(url, queryParamName) {
    document.getElementById("search-form").setAttribute("action", url);
    document.getElementById("search-form-query").setAttribute("name", queryParamName);
}

function runSearch() {
    document.querySelector("#search-form").submit();
}

document.addEventListener("DOMContentLoaded", function(event) {
    var popoverTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="popover"]'))
    popoverTriggerList.forEach(function (popoverTriggerEl) {
        const cid = popoverTriggerEl.getAttribute("popover-content-id")
        const content = document.querySelector('#' + cid)

        new bootstrap.Popover(popoverTriggerEl, {
            container: 'body',
            trigger: "focus",
            title: popoverTriggerEl.getAttribute("title"),
            html: true,
            content: content,
        })
    });
});

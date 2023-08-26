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

function showTurnOffContainerModal(containerId, containerName)
{
    let elem = document.getElementById('turnOffContainerModalContainer');

    elem.outerHTML = '<div class="modal fade" id="turnOffContainerModal" tabindex="-1" aria-labelledby="turnOffContainerModalLabel" aria-hidden="true">\n' +
        '    <div class="modal-dialog">\n' +
        '        <div class="modal-content">\n' +
        '            <div class="modal-header">\n' +
        '                <h5 class="modal-title" id="turnOffContainerModalLabel">Turn off ' + containerName + '?</h5>\n' +
        '                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>\n' +
        '            </div>\n' +
        '            <div class="modal-body" style="overflow: scroll;word-break: break-all;">\n' +
        '                You are about turn off container <strong>' + containerId + '</strong>?\n' +
        '            </div>\n' +
        '            <div class="modal-footer">\n' +
        '                <button type="button" class="btn btn-danger" onclick="window.location.href=\'/containers/kill?containerId=' + containerId + '\'">Turn off</button>\n' +
        '                <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>\n' +
        '            </div>\n' +
        '        </div>\n' +
        '    </div>\n' +
        '</div>'

    const modal = new bootstrap.Modal(document.getElementById('turnOffContainerModal'))

    modal.toggle()
}

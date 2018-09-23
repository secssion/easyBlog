$(document).ready(function () {
    var simplemde = new SimpleMDE({
        element: document.getElementById("editor"),
        status: false,
        lineNumbers: true,
        autoDownloadFontAwesome: false,
        tabSize: 4,
        renderingConfig: {
            codeSyntaxHighlighting: true
        },
    });
    simplemde.toggleSideBySide();
});
(function() {
    var editLink = document.querySelector('#edit-link a');

    if (editLink) {
        var editLinkP = document.getElementById('edit-link');
        var editForm = document.getElementById('edit-form');
        var cancelLink = document.querySelector('#edit-form p a');
        editLink.addEventListener('click', function(event) {
            event.preventDefault();
            editLinkP.className = 'hide';
            editForm.className = '';
        });
        cancelLink.addEventListener('click', function(event) {
            event.preventDefault();
            editForm.className = 'hide';
            editLinkP.className = '';
        });
    }
})();
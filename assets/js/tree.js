$(document).on('click', '#btn-exec', function(){
    $('#btn-exec').prop('disabled', true)
    let q = $('#q-input').val()
    let d = $('#d-input').val()
    let levels = $('#check-expand').prop('checked') ? d : 1
    $.ajax({
        type: 'GET',
        url: `/api/suggests/${encodeURIComponent(q)}/${d}`,
    }).done(function(data){
        $('#treeview').treeview({
            data: data,
            enableLinks: false,
            levels: levels,
            showTags: true,
            showBorder: false
        });
        let allLines = []
        let ser = function(node, depth=1) {
            allLines.push("\t".repeat(depth-1)+node.text)
            if ('nodes' in node) {
                for (let child of node.nodes) {
                    ser(child, depth+1)
                }
            }
        }
        for (let d of data) {
            ser(d)
        }
        $('#treeviewtext').empty()
        $('#treeviewtext').append('<textarea class="form-control" rows="20">'+allLines.join("\n")+"</textarea>")
        $('#filter-box').show()
        $('#filter-result').empty()
        $('#filter-cap').text('')
        $('#result-container').show()
    }).fail(function(e){
      alert("Request failed! Try again later.")
    }).always(function(){
        $('#btn-exec').prop('disabled', false)
    });
});

$(document).on('nodeSelected nodeExpanded', function(e, node) {
    $('#filter-result').empty()
    if ('nodes' in node) {
        let res = node.nodes.map(x => `<li class="list-group-item">${x.text}</li>`)
        for (let r of res) {
            $('#filter-result').append(r)
        }
    }
    $('#filter-cap').text(`${node.text}`)
})

$(document).on('keypress', '#q-input', function(e) {
    if(e.which == 13) {
        $('#btn-exec').click()
    }
})

$(document).on('change', '#check-expand', function() {
    if ($(this).prop('checked')) {
        $('#treeview').treeview('expandAll', { levels: 3, silent: true });
    } else {
        $('#treeview').treeview('collapseAll', { silent: true });
    }
})
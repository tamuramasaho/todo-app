$(function(){
    $('#create').submit(function(e){
      e.preventDefault();
      var $form = $(this);
      var $button = $form.find('button');
      var data = $('#create').serializeArray();
      data = parseJson(data)
      $.ajax({
        url:           $form.attr('action'),
        type:          $form.attr('method'),
        dataType:      'json',
        contentType:   'application/json',
        scriptCharset: 'utf-8',
        data:          JSON.stringify(data),
        timeout: 10000,
        beforeSend: function(xhr, settings) {
          $button.attr('disabled', true);
        }
      })
      .done(function(result) {
        $button.attr('disabled', false);
        $form[0].reset();
        $('#table-todos').append(htmlTodos(result));
      }).fail(alertError);
    });

    $('#update').submit(function(e){
      e.preventDefault();
      var $form = $(this);
      var data = $('#update').serializeArray();
      data = parseJson(data)
      $.ajax({
        url:           $form.attr('action'),
        type:          $form.attr('method'),
        dataType:      'json',
        contentType:   'application/json',
        scriptCharset: 'utf-8',
        data:          JSON.stringify(data),
        timeout: 10000,
      })
      .done(function(result) {
        window.location.href = '/todos';
      }).fail(alertError);
    });

    $(document).on("click",'.delete', function(){
      var $button = $(this);
      $.ajax({
        url:           $button.attr('action'),
        type:          $button.attr('method'),
        scriptCharset: 'utf-8',
      })
      .done(function(result) {
        $button.parent().parent().hide();
      }).fail(alertError);
    });

    $(document).on("click",'.unfinished', function(){
      var $button = $(this);
      $.ajax({
        url:           $button.attr('action'),
        type:          $button.attr('method'),
        scriptCharset: 'utf-8',
      })
      .done(function(result) {
        $button.parent().parent().hide();
        $('#table-done').append(htmlTodos(result));
        $('#table-done').find('.unfinished').removeClass( "unfinished" ).addClass( "finished" );
      }).fail(alertError);
    });

    $(document).on("click",'.finished', function(){
      var $button = $(this);
      $.ajax({
        url:           $button.attr('action'),
        type:          $button.attr('method'),
        scriptCharset: 'utf-8',
      })
      .done(function(result) {
        $button.parent().parent().hide();
        $('#table-todos').append(htmlTodos(result));

      }).fail(alertError);
    });

    $(document).on("click",'.reminder', function(){
      var $button = $(this);
      $.ajax({
        url:           $button.attr('action'),
        type:          $button.attr('method'),
        scriptCharset: 'utf-8',
        beforeSend: function(xhr, settings) {
          $button.attr('disabled', true);
        },
      })
      .done(function(result) {
        $button.attr('disabled', false);
        $button.next().text(result["RemindAt"])
      }).fail(alertError);
    });

    var alertError = function(result) {
      alert("エラーが発生しました。");
      window.location.href = '/todos';
    }

    var htmlTodos = function(result) {
    return `<tr>
      <td>
        <button class="fas fa-times cursor unfinished" 
              action="/todos/${result["ID"]}" method="GET"></button>
      </td>
      <td class="have" style="table-layout:fixed;width:30%;">${result["Title"]}</td>
      <td class="have" style="table-layout:fixed;width:10%;">${result["DueDate"]}</td>
      <td class="have" style="table-layout:fixed;width:35%;">
        <button class="reminder cursor"
                        action="/todos/${result["ID"]}" method="PATCH"><i class="fas fa-bell fa-2x"></i></button>
        <span>${result["RemindAt"]}</span>
      </td>
      <td style="table-layout:fixed;width:35%;">
          <a href="/todos/${result["ID"]}/edit"><i class="fas fa-edit fa-2x"></i></a>
          <button class="delete cursor fas fa-trash-alt fa-2x" 
                        action="/todos/${result["ID"]}" method="DELETE"></button>
      </td>
    </tr>`
    }

    var parseJson = function(data) {
        var returnJson= {};
        for (idx = 0; idx < data.length; idx++) {
          returnJson[data[idx].name] = data[idx].value
        }
        return returnJson;
    }
});


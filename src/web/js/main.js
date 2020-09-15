//Good jquery practise for encapsulation
$(document).ready(function(){

  /*
  Method bound to the input event of the #primeInput element.
  Sends an ajax request to api/GetPreviousPrime with the number provided by the user
  */
  function OnInputChange(e){
    var currentValue = e.target.value;
    $.ajax({
      url: "api/GetPreviousPrime",
      data: {number: currentValue},
      type: "GET",
      success: UpdateResult,
      dataType: "json"
    });
  }

  /*
  Callback for the OnInputChange ajax request
  Shows any and all of the result elements related to the operation
  */
  function UpdateResult(e){
    if(e.previousPrime == 0){
      $("#result").hide();
      $("#error").show();
      $("#success").hide();
      return;
    }
    $("#error").hide();
    $("#success").show();
    $("#result").text(e.previousPrime).show();
  }

  $("#primeInput").on("input", OnInputChange);
})

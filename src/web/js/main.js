$(document).ready(function(){

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

$(function() {
  $("#minus-btn").click(function(){
    var amount = parseInt($("#amount").val());
    if (amount > 1) amount -= 1;
    $("#amount").val(amount);
    $("#pay-amount").text("10元 X "+amount+" = "+(amount*10)+" 元");
  });

  $("#plus-btn").click(function(){
    var amount = parseInt($("#amount").val())+1;
    $("#amount").val(amount);
    $("#pay-amount").text("10元 X "+amount+" = "+(amount*10)+" 元");
  });

  $("#pay-btn").click(function() {
    pingpp_one.init({
      app_id:"", // from auth
      order_no:"", // from auth
      amount:parseInt($("#amount").val()),
      channel:["alipay_wap"],
      charge_url:"/pay"
    }, function(res) {
      if (!res.status) _toast.show(res.msg);
    });
  });
});

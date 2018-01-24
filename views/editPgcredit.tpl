<!DOCTYPE html>
<html>
	<head>
		{{template "header"}}
		<script src="http://cdn.static.runoob.com/libs/jquery/2.1.1/jquery.min.js"></script>
		<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
	</head>
	<body>
		{{template "TopBar"}}
		<div class="container-fluid" style="padding-top:50px">
      		<div class="row">
        		{{template "LeftBar"}}
				<div class="col-sm-9">
					<div class="panel panel-primary">
						<div class="panel-heading">
							<h3 class="panel-title">查看专业毕业学分</h3>
						</div>
					    <div class="panel-body">
							<label for="name">专业：{{.pmname}}</label>
							{{range .pgc_info}}
							<table border="1" cellpadding="1" cellspacing="1" class="table-bordered table-condensed" style="width: 700px; height: 180px;">
								<tbody>
									<tr>
										<td colspan="1" rowspan="3" style="width:50px">理论课程</td>
										<td colspan="2" rowspan="1">公共</td>
										<td colspan="3" rowspan="1">专业</td>
										<td>小计</td>
										<td>总学分</td>
									</tr>
									<tr>
										<td>必修</td>
										<td>任选</td>
										<td>必修</td>
										<td>限选</td>
										<td>任选</td>
										<td>&nbsp;</td>
										<td>&nbsp;</td>
									</tr>									
									<tr>
										<td><input id="ggbx" type="text" class="form-control" value="{{.Open_require_credit}}"></td>
										<td><input id="ggrx" type="text" class="form-control" value="{{.Open_option_credit}}"></td>
										<td><input id="zybx" type="text" class="form-control" value="{{.Professional_require_credit}}"></td>
										<td><input id="zyxx" type="text" class="form-control" value="{{.Professional_option_credit}}"></td>
										<td><input id="zyrx" type="text" class="form-control" value="{{.Professional_limit_credit}}"></td>
										<td><input id="llxj" type="text" class="form-control"></td>
										<td><input id="zxf" type="text" class="form-control" value="{{.Total_credit}}"></td>
									</tr>
									<tr>
										<td>实际环节</td>
										<td colspan="5" rowspan="1"><input id="sjxf" type="text" class="form-control" value="{{.Practice_credit}}"></td>
										<td><input id="sjxj" type="text" class="form-control"></td>
										<td></td>
									</tr>
									
								</tbody>
							</table>
							<div class="col-sm-3 pull-right">
								<button type="button" class="btn btn-primary" onclick="return SumInput()">重新计算</button>
								<button type="button" class="btn btn-primary" onclick="return UpdateInput()">更新</button>
								<button type="button" class="btn btn-primary" onclick="return DelInput()">删除</button>
							</div>
						</div>
					{{end}}	
					</div>
				</div>
			</div>			
		</div>
		<script type="text/javascript">
			function  UpdateInput(){
				//alert("点击更新按钮")	
				var llxiaoji=parseFloat($("#ggbx").val())+parseFloat($("#ggrx").val())+parseFloat($("#zybx").val())+parseFloat($("#zyxx").val())+parseFloat($("#zyrx").val())
				var sjxiaoji=parseFloat($("#sjxf").val())
				var zxf=llxiaoji+sjxiaoji		
				$.ajax({
					type:"POST",
					url:"{{urlfor "PGCreditController.PgcUpdate"}}",
					data:{
						pmname:{{.pmname}},
						ggbx:parseFloat($("#ggbx").val()),
						ggrx:parseFloat($("#ggrx").val()),
						zybx:parseFloat($("#zybx").val()),
						zyxx:parseFloat($("#zyxx").val()),
						zyrx:parseFloat($("#zyrx").val()),
						sjxf:parseFloat($("#sjxf").val()),
						zxf:zxf			
					},
					async:false,
					error:function(request){
						alert("post error")						
					},
					success:function(data){
						if(data.status==0){
							alert("更新成功")
							window.location.href="/pgcredit/edit?pmname="+{{.pmname}}
						}else{
							alert("更新失败")
						}
						
					}
					
				});
				
				return true	
			}

		function SumInput(){
			//alert("点击重新计算")
			var llxiaoji=parseFloat($("#ggbx").val())+parseFloat($("#ggrx").val())+parseFloat($("#zybx").val())+parseFloat($("#zyxx").val())+parseFloat($("#zyrx").val())
				var sjxiaoji=parseFloat($("#sjxf").val())
				var zxf=llxiaoji+sjxiaoji				
				$("#llxj").val(llxiaoji)
				$("#sjxj").val(sjxiaoji)
				$("#zxf").val(zxf)
		}
		function DelInput(){
			//alert("点击删除")
				$.ajax({
					type:"POST",
					url:"{{urlfor "PGCreditController.PgcDel"}}",
					data:{pmname:{{.pmname}}},
					async:false,
					error:function(request){
						alert("post error")				
					},
					success:function(data){
						if(data.status==0){
							alert("删除成功")
							window.location.href="/pgcredit"
						}else{
							alert("删除失败")
						}						
					}
					
				});
				return true	
		}	
		</script>
	</body>		
</html>

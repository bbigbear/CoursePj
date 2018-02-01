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
							<h3 class="panel-title">设置专业毕业学分</h3>
						</div>
					    <div class="panel-body">
							<label for="name">专业：{{.pmname}}</label>
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
										<td><input id="ggbx"  class="form-control" value="0.0" type="number"></td>
										<td><input id="ggrx"  class="form-control" value="0.0" type="number"></td>
										<td><input id="zybx"  class="form-control" value="0.0" type="number"></td>
										<td><input id="zyxx"  class="form-control" value="0.0" type="number"></td>
										<td><input id="zyrx"  class="form-control" value="0.0" type="number"></td>
										<td><input id="llxj"  class="form-control" ></td>
										<td><input id="zxf" class="form-control" ></td>
									</tr>
									<tr>
										<td>实际环节</td>
										<td colspan="5" rowspan="1"><input id="sjxf"  class="form-control" value="0.0" type="number"></td>
										<td><input id="sjxj"  class="form-control"></td>
										<td></td>
									</tr>
								</tbody>
							</table>
							<div class="col-sm-2 pull-right">
								<button type="button" class="btn btn-primary" onclick="return SumInput()">计算</button>
								<button type="button" class="btn btn-primary" onclick="return SaveInput()">保存</button>
							</div>
						</div>						
					</div>
				</div>
			</div>			
		</div>
		<script type="text/javascript">
			function SumInput(){
				//alert("点击计算")
				var llxiaoji=parseFloat($("#ggbx").val())+parseFloat($("#ggrx").val())+parseFloat($("#zybx").val())+parseFloat($("#zyxx").val())+parseFloat($("#zyrx").val())
				var sjxiaoji=parseFloat($("#sjxf").val())
				var zxf=llxiaoji+sjxiaoji
				
				$("#llxj").val(llxiaoji)
				$("#sjxj").val(sjxiaoji)
				$("#zxf").val(zxf)
			}
			function SaveInput(){
				//alert("点击保存")
				var llxiaoji=parseFloat($("#ggbx").val())+parseFloat($("#ggrx").val())+parseFloat($("#zybx").val())+parseFloat($("#zyxx").val())+parseFloat($("#zyrx").val())
				var sjxiaoji=parseFloat($("#sjxf").val())
				var zxf=llxiaoji+sjxiaoji
				$.ajax({
					type:"POST",
					url:"/pgcredit/save",
					data:{
						pmname:{{.pmname}},
						year:{{.y}},
						faculty:{{.f}},
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
						if(data.status==200){
							alert("保存成功")
							window.location.href="/pgcredit/search?year="+{{.y}}+"&faculty="+{{.f}}
						}else{
							alert("保存失败")
						}
						
					}
					
				});
			}
			function AddInput(){
				var cid=document.getElementById("pmid")
				if(pmid.value.length==0){
					alert("专业代码不能为空")
					return false
				}
				
				$.ajax({
					type:"POST",
					url:"{{urlfor "PmController.PmAddAction"}}",
					data:$("#addPm").serialize(),
					async:false,
					error:function(request){
						alert("post error")
						
					},
					success:function(data){
						if(data.status==200){
							alert("新增成功")
							window.location.href="/pm/add"
						}else{
							alert("新增失败")
						}
						
					}
					
				});
				
				return true	
			}	
		</script>
	</body>
		
</html>

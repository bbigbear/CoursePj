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
							<label for="name">专业：</label>
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
										<td><input type="text" class="form-control"></td>
										<td><input type="text" class="form-control"></td>
										<td><input type="text" class="form-control"></td>
										<td><input type="text" class="form-control"></td>
										<td><input type="text" class="form-control"></td>
										<td><input type="text" class="form-control"></td>
										<td><input type="text" class="form-control"></td>
									</tr>
									<tr>
										<td>实际环节</td>
										<td colspan="5" rowspan="1"><input type="text" class="form-control"></td>
										<td><input type="text" class="form-control"></td>
										<td><input type="text" class="form-control"></td>
									</tr>
								</tbody>
							</table>
							<div class="col-sm-2 pull-right">
								<button type="button" class="btn btn-primary" onclick="return EditInput()">修改</button>
								<button type="button" class="btn btn-primary" onclick="return EditInput()">删除</button>
							</div>
						</div>
						
					</div>
				</div>
			</div>			
		</div>
		<script type="text/javascript">
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
						if(data.status==0){
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

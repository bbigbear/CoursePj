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
				<div class="col-sm-5">
					<div class="panel panel-primary">
						<div class="panel-heading">
							<h4 class="panel-title">未设置专业毕业学分</h4>
						</div>
					    <div class="panel-body">
						<div class="row">				
					    <div class="col-sm-6">							
							<form role="form">
							  <div class="form-group">
							    <label for="name">未设置专业学分列表：{{.Pmslice_NotSet_count}}个</label>
							    <select multiple class="form-control" id="Pmslice_NotSet">
								{{range .Pmslice_NotSet}}	
							      <option>{{.}}</option>							   
								{{end}}	
							    </select>
							  </div>
							</form>																			
						</div>
						<div class="col-sm-2" style="padding-top:25px">
							<button type="button" class="btn btn-primary" onclick="return SetInput()">设置专业学分</button>
						</div>
						</div>						
						</div>
					</div>					
				</div>
				<div class="col-sm-5">
					<div class="panel panel-primary">
						<div class="panel-heading">
							<h4 class="panel-title">已设置专业毕业学分</h4>
						</div>
						<div class="panel-body">
						<div class="row">				
						    <div class="col-sm-6">							
								<form role="form">
								  <div class="form-group">
								    <label for="name">已设置专业学分列表：{{.Pmslice_Set_count}}个</label>
								    <select multiple class="form-control" id="Pmslice_Set">
									{{range .Pmslice_Set}}	
								      <option>{{.}}</option>							   
									{{end}}	
								    </select>
								  </div>
								</form>																			
							</div>
						<div class="col-sm-2" style="padding-top:25px">
							<button type="button" class="btn btn-primary" onclick="return EditInput()" style="margin-top:10px">查看专业学分</button>
						</div>
					    </div>
						
						</div>
					</div>							
				</div>
      		</div>
    	</div>
		<script type="text/javascript">
			
			function EditInput(){				
				var pmname=document.getElementById("Pmslice_Set")
				//alert($("#Pmid").is(":checked"))
				//alert($("#Pmslice_Set").val())
				if ($("#Pmslice_Set").val()!=null){
					window.location.href="/pgcredit/edit?pmname="+pmname.value
				}else{
					alert("请选择专业再查看学分")
				}				
			}
			function SetInput(){				
				var pmname=document.getElementById("Pmslice_NotSet")
				//alert($("#Pmid").is(":checked"))
				//alert($("#Pmslice_Set").val())
				if ($("#Pmslice_NotSet").val()!=null){
					window.location.href="/pgcredit/add?pmname="+pmname.value
				}else{
					alert("请选择专业再设置学分")
				}
							
							
			}		
		</script>
	</body>
		
</html>

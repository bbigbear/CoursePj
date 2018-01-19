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
				<div class="col-sm-10">
					<div class="panel panel-primary">
						<div class="panel-heading">
							<h4 class="panel-title">查看专业理论课程</h4>
						</div>
					    <div class="panel-body">
						<div class="row">				
					    <div class="col-sm-4">							
							<form role="form">
							  <div class="form-group">
							    <label for="name" id="Pname">{{.pmid}}</label>
							    <select multiple class="form-control" id="Cname">
								{{range .s}}
							      <option>{{.}}</option>		   
								{{end}}	
							    </select>
							  </div>
							</form>														
						</div>
						<div class="col-sm-2" style="padding-top:25px">
							<button type="button" class="btn btn-primary" onclick="return ToDelCourse()">删除选中课程</button>
						</div>
						</div>
						</div>
					</div>
				</div>
      		</div>
    	</div>
		<script type="text/javascript">
			
			function ToDelCourse(){											
				//var pname=document.getElementById("Pname")
				var data="";
				$("#Cname :selected").each(function(){				 
					 data=data+$(this).val()+',';
				 });
				//alert(data)
				$.ajax({
					type:"POST",
					url:"{{urlfor "PTCourseController.PTCourseDelete"}}",
					data:{cname:data,pmid:{{.pmid}}},
					async:false,
					error:function(request){
						alert("post error")		
					},
					success:function(data){
						if(data.status==0){
							alert("删除成功")
							window.location.href="/ptcourse/edit?pmid="+{{.pmid}}
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

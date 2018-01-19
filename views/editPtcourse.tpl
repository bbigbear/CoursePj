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
							    <label for="name">{{.pmid}}</label>
							    <select multiple class="form-control" id="Cid">
								{{range .m}}	
							      <option>{{.Cid}}</option>		   
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
			
			function QueryInput(){				
				var s_Status=document.getElementById("s_Status")
				var s_Year=document.getElementById("s_Year")
				window.location.href="/home/search?s_Year="+s_Year.value+"&s_Status="+s_Status.value
			}
			function ToDelCourse(){
				var checked_array=[];
				var selectedValues=[];
				var data="";
				var Pm_data="";
			 	$("[name='Cid']:checkbox:checked").each(function(){				 
					checked_array.push($(this).val()) 	
					data=data+$(this).val()+',';			
				});
				$("#Pmid :selected").each(function(){
				     selectedValues.push($(this).val());
					 Pm_data=Pm_data+$(this).val()+',';
				 });
				//alert(Pm_data)
				//alert(data)				
				$.ajax({  
				    url: "{{urlfor "PTCourseController.Setcourse"}}",  
				    data: { cid: data,pmid: Pm_data},    
				    type: "POST",
					async:false,
					error:function(data){
						alert("post error")
					},
				    success:function(data){  
				        if(data.status==0){
							alert("设置成功")
						}else{
							alert("设置失败，已存在相关理论课程")
						}
				    }  
				});			
			}		
		</script>
	</body>
		
</html>

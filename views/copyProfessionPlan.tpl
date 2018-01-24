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
							<h4 class="panel-title">查询专业信息</h4>
						</div>
					    <div class="panel-body">
						<form class="form-inline" role="form" id="searchPm">
					        <div class="form-group">				
								<label>年级</label>			
								<select class="form-control" name="s_Pmyear" id="s_Pmyear">
								<option>2015</option>
								<option>2016</option>
								<option>2017</option>
								<option>2018</option>
								</select>
								<label>院系</label>			
								<input class="form-control" name="s_Pmfaculty" id="s_Pmfaculty">
								<button type="button" class="btn btn-default" onclick="return QueryInput()">检索</button>				
							</div>
						</form>
						<div class="col-sm-1 pull-right">					
							<button type="button" class="btn btn-primary" onclick="return AddInput()">复制</button>																	 	
					    </div>						
						</div>																	
					</div>					
				</div>				
      			<div class="col-sm-5">
					<div class="panel panel-primary">
					    <div class="panel-body">
						<div class="row">				
					    <div class="col-sm-8">							
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
						</div>						
						</div>
					</div>					
				</div>
				<div class="col-sm-5">
					<div class="panel panel-primary">
						<div class="panel-body">
						<div class="row">				
						    <div class="col-sm-8">							
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
						</div>						
					</div>
				</div>							
				</div>
			</div>			
    	</div>
		<script type="text/javascript">
			
			function QueryInput(){
				var s_Pmyear=document.getElementById("s_Pmyear")
				var s_Pmfaculty=document.getElementById("s_Pmfaculty")
				window.location.href="/pm/search?s_Pmyear="+s_Pmyear.value+"&s_Pmfaculty="+s_Pmfaculty.value
			}
			function AddInput(){
				window.location.href="/pm/add"
			}
			
			var isCheckAll=false;
			function swapCheck(){
				if(isCheckAll){
					$("input[name='pm1id']").each(function(){
						this.checked=false
					});
					isCheckAll=false;
				}else{
					$("input[name='pm1id']").each(function(){
						this.checked=true
					});
					isCheckAll=true;
				}
			}
			function ToAble(){
				var checked_array=[];
				var data="";
			 	$("[name='pm1id']:checkbox:checked").each(function(){				 
					checked_array.push($(this).val()) 	
					data=data+$(this).val()+',';			
				});
				alert(data)
				$.ajax({  
				    url: "{{urlfor "PmController.PmStautsChange"}}",  
				    data: { pm1id: data},    
				    type: "POST",
				    success: function () {  
				        // your logic 
				        alert('Ok');  
				    }  
				});			
			}
		</script>
	</body>
		
</html>

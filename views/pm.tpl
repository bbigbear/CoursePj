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
								<select class="form-control" name="s_Pmfaculty" id="s_Pmfaculty">
								{{range .f}}
								<option>{{.}}</option>
								{{end}}				
								</select>		
<!--								<input class="form-control" name="s_Pmfaculty" id="s_Pmfaculty">-->
								<button type="button" class="btn btn-default" onclick="return QueryInput()">检索</button>				
							</div>
						</form>
						<div class="col-sm-1 pull-right">					
							<button type="button" class="btn btn-primary" onclick="return AddInput()">新增</button>																	 	
					    </div>
						</div>
					</div>
					<div class="row">
							<div class="col-sm-6">
							<div class="panel panel-default" >
								<div class="panel-heading">
									<h5 class="panel-title">未开设专业</h5>								
								</div>
								<div class="panel-body">
									<div>
									<label><input type="checkbox" onclick="return swapCheck()">全选</label>
									<button class="pull-right" onclick="return ToAble()">-></button>
									</div>
									<div class="table-responsive">
										<table class="table table-bordered">											
											<thead>
												<tr>
													<th>选择</th>
													<th>操作</th>
													<th>院系</th>
													<th>专业代码</th>
													<th>专业名称</th>
													<th>培养层次</th>
													<th>允许辅修</th>
													<th>英文名称</th>
													<th>状态</th>
												</tr>
												{{range .m1}}		
												<tr>
													<th><input type="checkbox" value="{{.Pmid}}" name="pm1id"></th>
													<th><a href="/pm/edit?pmid={{.Pmid}}&year={{.Year}}">编辑</a></th>
													<th>{{.Faculty}}</th>
													<th>{{.Pmid}}</th>
													<th>{{.Pmname}}</th>
													<th>{{.Train_level}}</th>
													<th>{{.Isminor}}</th>
													<th>{{.Pmname_en}}</th>
													<th>{{.Status}}</th>							
												</tr>
												{{end}}											
											</thead>
											<tbody>							
											</tbody>
										</table>
									</div>
								</div>
							</div>							
							</div>
							<div class="col-sm-6">
							<div class="panel panel-default">
								<div class="panel-heading">
									<h3 class="panel-title">已开设专业</h3>
								</div>
								<div class="panel-body">
									<div class="row">
									<button class="pull-left"><-</button>
									</div>
									<div class="table-responsive">
										<table class="table table-bordered">											
											<thead>
												<tr>
													<th>选择</th>
													<th>操作</th>
													<th>院系</th>
													<th>专业代码</th>
													<th>专业名称</th>
													<th>培养层次</th>
													<th>允许辅修</th>
													<th>英文名称</th>
													<th>状态</th>
												</tr>
												{{range .m}}
												<tr>
													<th><input type="checkbox" value="{{.pmid}}"></th>
													<th><a href="/pm/edit?pmid={{.Pmid}}&year={{.Year}}">编辑</a></th>
													<th>{{.Faculty}}</th>
													<th>{{.Pmid}}</th>
													<th>{{.Pmname}}</th>
													<th>{{.Train_level}}</th>
													<th>{{.Isminor}}</th>
													<th>{{.Pmname_en}}</th>
													<th>{{.Status}}</th>								
												</tr>
												{{end}}											
											</thead>
											<tbody>							
											</tbody>
										</table>
									</div>
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

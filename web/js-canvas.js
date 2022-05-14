fetch('http://localhost:8080/route?fromCache=true', {method: 'GET', mode: 'cors', cache: 'no-cache'})
.then(response => {
   return response.json();
})
.then(routesdata => {
	var canvas = document.getElementById("main-water");
	var x_size = routesdata["TableSize"]["X"];
	var y_size = routesdata["TableSize"]["Y"];
	canvas.width = x_size;
	canvas.height = y_size;
	
	var ctx = canvas.getContext("2d");
	for (route_index = 0; route_index < routesdata["Routes"].length; route_index++) {
		var route_steps = routesdata["Routes"][route_index]["Steps"];
		var route_step_count = route_steps.length;
		ctx.beginPath();
		
		ctx.fillStyle = "#9a6642";
		ctx.fillRect(route_steps[0]["X"], route_steps["Y"], 5, 5);
		
		ctx.lineWidth = 1;
		var random_color = (Math.random() * 0xFFFFFF << 0).toString(16).padStart(6, '0');
		ctx.strokeStyle = "#" + random_color; 
		ctx.moveTo(route_steps[0]["X"], route_steps[0]["Y"]);
		
		for (step_index = 1; step_index < route_step_count; step_index++) {
			ctx.lineTo(route_steps[step_index]["X"], route_steps[step_index]["Y"]);
		}
		ctx.stroke();
		
		ctx.fillStyle = "#9a6642";
		ctx.fillRect(route_steps[route_step_count - 1]["X"], route_steps[route_step_count - 1]["Y"], 5, 5);
	}
});

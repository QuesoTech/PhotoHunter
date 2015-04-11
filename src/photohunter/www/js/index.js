/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

var pictureSource;   // picture source
var destinationType; // sets the format of returned value

document.addEventListener('deviceready', onDeviceReady, false);

function onDeviceReady() {
		pictureSource = navigator.camera.PictureSourceType;
		destinationType=navigator.camera.DestinationType;
}


// Called when a photo is successfully retrieved
//
function onPhotoDataSuccess(imageData) {
	// Uncomment to view the base64-encoded image data
	// console.log(imageData);

	// Get image handle
	//
	var smallImage = document.getElementById('smallImage');

	// Unhide image elements
	//
	smallImage.style.display = 'block';

	// Show the captured photo
	// The in-line CSS rules are used to resize the image
	//
	smallImage.src = "data:image/jpeg;base64," + imageData;
}

var onGeoSuccess = function(position) {
		$('#coords').html(String(position.coords.latitude) + ", " + String(position.coords.longitude))	
		$('#coords').attr('class', 'label label-success')	
};

function capturePhoto()
{
	navigator.camera.getPicture(onPhotoDataSuccess, onFail, { quality: 50, 
		destinationType: destinationType.DATA_URL });
	getLocation();
	$('#takepicture').attr('style', 'display:none;')
	$('#upload').attr('style', 'display:block; margin:0 auto;')
}

function getLocation()
{
	navigator.geolocation.getCurrentPosition(onGeoSuccess, onFail);
}

// Called if something bad happens.
//
function onFail(message) {
	alert('Failed because: ' + message);
}


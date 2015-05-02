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
var userID;

document.addEventListener('deviceready', onDeviceReady, false);

function onDeviceReady() {
		pictureSource = navigator.camera.PictureSourceType;
		destinationType=navigator.camera.DestinationType;
		
		// Check if the account button is on the current page
		// and sets the image to be the user's profile picture
		if (document.getElementById('accountButton'))
		{
			getUserProfilePic("small", "accountButton");
		}
		if (document.getElementById('accountPic'))
		{
			getUserProfilePic("large", "accountPic");
			fillAccountInfo();
		}
}

//
// Called when a photo is successfully retrieved
//
function onPhotoDataSuccess(imageData) {
	// Uncomment to view the base64-encoded image data
	// console.log(imageData);

	// Get image handle
	var smallImage = document.getElementById('smallImage');

	// Unhide image elements
	smallImage.style.display = 'block';

	// Show the captured photo
	// The in-line CSS rules are used to resize the image
	smallImage.src = "data:image/jpeg;base64," + imageData;
}


//
// Called when the geolocation is successfully retrieved
//
var onGeoSuccess = function(position) {
		// Display the coordinates and turn the labels green
		$('#coords').html(String(position.coords.latitude) + ", " + String(position.coords.longitude))	
		$('#coords').attr('class', 'label label-success')	
};


//
// Called when a user successfully logs in to the app
//
var fbLoginSuccess = function (userData) {
    facebookConnectPlugin.getAccessToken(function(token) 
		{
				// Code to grab datasets will go here

				window.location.replace("datasets.html");
    }, 
		function(err) 
		{
			alert("Could not get access token: " + err);
    });
};


//
// Open the native camera app and take a picture, then display the
// geolocation and upload button
//
function capturePhoto()
{
	navigator.camera.getPicture(onPhotoDataSuccess, onFail, { quality: 50, 
		destinationType: destinationType.DATA_URL });
	getLocation();
	$('#takepicture').attr('style', 'display:none;')
	$('#upload').attr('style', 'display:block; margin:0 auto;')
}

//
// Login through Facebook
//
function loginFB()
{

	facebookConnectPlugin.login(["public_profile"],	fbLoginSuccess,	onFail);

}

//
//	Get the profile picture of the signed in
//  user and display it as the account page
//  button
//
function getUserProfilePic(size, element)
{
	facebookConnectPlugin.getLoginStatus( 
		function (response)
		{
			// Get user id and change the account button image
			userID = response.authResponse["userID"];
			var accountPic = document.getElementById(element);
			accountPic.src = "https://graph.facebook.com/" + userID + "/picture?type=".concat(size)
		},
		function (error){alert("Failed: " + error)});
}

//
// Fill out the info on the account page
//
function fillAccountInfo()
{
	facebookConnectPlugin.getLoginStatus( 
		function (response)
		{
			// Get user id and change the account button image
			userID = response.authResponse["userID"];
			facebookConnectPlugin.api(userID, ["public_profile"],
				function (response)
				{
					userName = response["name"];
					accountName.innerText = userName;
				},
				function (error){alert("Failed: " + error)});
		},
		function (error){alert("Failed: " + error)});
}


//
// Get the user's current location
//
function getLocation()
{
	navigator.geolocation.getCurrentPosition(onGeoSuccess, onFail);
}

//
// Called if something bad happens.
//
function onFail(message) {
	alert('Failed because: ' + message);
}


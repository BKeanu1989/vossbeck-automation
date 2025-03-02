console.log("hello world");
(function () {
  if (!window.Trello) {
    console.warn("trello not defined");
    return;
  }

  var authenticationSuccess = function () {
    console.log("Successful authentication");
  };

  var authenticationFailure = function () {
    console.log("Failed authentication");
  };

  window.Trello.authorize({
    type: "popup",
    name: "Getting Started Application",
    scope: {
      read: "true",
      write: "true",
    },
    expiration: "never",
    success: authenticationSuccess,
    error: authenticationFailure,
  });
})();

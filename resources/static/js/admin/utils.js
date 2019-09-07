var storageAuthTokenKey = 'token';

/**
 *
 */
function redirect(url, needAuth = true) {
    if (needAuth) {
        url += "?token"
    }
    //判断是否有token
    location.href = url;
}


//保存token
function storageAuthToken(token) {
    localStorage.setItem(storageAuthTokenKey, token)
}

//获取token
function getAuthToken() {
    return localStorage.getItem(storageAuthTokenKey)
}
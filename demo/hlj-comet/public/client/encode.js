function buildQuery(param, key, encode) {
    if (param == null) return '';

    let t = typeof (param);
    if (t == 'string' || t == 'number' || t == 'boolean') {
        return key + '=' + ((encode == null || encode) ? encodeURIComponent(param) : param);

    }

    let paramStr = '';
    for (let i in param) {
        let k = key == null ? i : key + (param instanceof Array ? '[' + i + ']' : '.' + i);
        let delmiter = paramStr === '' ? '' : '&';
        paramStr += delmiter + buildQuery(param[i], k, encode);
    }
    return paramStr;
}
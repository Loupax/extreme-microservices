function jwt(data) {
    var parts = data.split('.').slice(0,2)
    .map(v=>Buffer.from(v, 'base64url').toString())
    .map(JSON.parse);
    return { headers:parts[0], payload: parts[1] };
}
    
function jwt_claim_sub(r) {
    return jwt(r.headersIn.Authorization.slice(7)).payload.sub;
}
export default {jwt_claim_sub}
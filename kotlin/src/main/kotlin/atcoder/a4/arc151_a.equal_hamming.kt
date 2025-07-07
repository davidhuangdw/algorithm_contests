package atcoder.a4

fun main(){
    readLine()
    val S = readLine()!!
    val T = readLine()!!
    val n = S.length
    val diff = (0 until n).filter{ S[it] != T[it]}
    if(diff.size % 2 != 0){
        println(-1)
        return
    }

    val k = diff.size /2
    val res = CharArray(n){'0'}
    var (a, b) = 0 to 0
    for(i in diff){
        if(a < k && b < k){
            if(S[i] == '0'){
                a += 1
            }else {
                b += 1
            }
        }else if(a >= k){
            res[i] = T[i]
        }else{
            res[i] = S[i]
        }
    }
    println(res.joinToString(""))
}
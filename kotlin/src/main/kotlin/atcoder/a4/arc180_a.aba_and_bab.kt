package atcoder.a4

fun main(){
    readLine()
    val S = readLine()!!
    val N = S.length
    val MOD = (1e9+7).toInt()
    var sum = 1L
    var i = 0
    while(i < N){
        val j = (i+1 until N).find{j -> S[j] == S[j-1]} ?: N
        sum = sum * (1+(j-1-i)/2) % MOD
        i = j
    }
    println(sum)
}
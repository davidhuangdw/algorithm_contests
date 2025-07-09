package atcoder.a2

fun main(){
    readLine()
    val s = readLine()!!
    val n = s.length
    var i = 0
    var sum = 0L
    while(i < n){
        val k = ((i+1 until n).find{ s[it] != s[i]} ?: n) - i
        sum += (1L * k * (k-1)) shr 1
        i += k
    }
    println(sum)
}
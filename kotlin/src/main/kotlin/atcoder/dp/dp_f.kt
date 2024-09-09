package atcoder.dp

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }

    val S = readLine()!!
    val T = readLine()!!
    val n = S.length
    val m = T.length
    val dp = Array(n+1) { IntArray(m+1) { 0 } }
    for(i in 1..n)
        for(j in 1..m)
            dp[i][j] = maxOf(dp[i-1][j], dp[i][j-1], if(S[i-1] == T[j-1]) dp[i-1][j-1]+1 else 0)

    var (i, j) = n to m
    val co = mutableListOf<Char>()
    while(i > 0 && j > 0 && dp[i][j] > 0){
        if(S[i-1] == T[j-1]) { // bug: need to ensure equal
            co.add(S[i-1])
            i--
            j--
        }else if(dp[i][j] == dp[i-1][j]){
            i--
        }else{
            j--
        }
    }
    println(co.reversed().joinToString(""))
}

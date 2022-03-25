package google_kickstart.`2017`.A

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    fun getPre(s: String): List<Int>{
        val n = s.length
        val pre = MutableList(n+1){0}
        for((i, x) in s.withIndex()){
            pre[i+1] = if(x == '*') pre[i] else i+1
        }
        return pre
    }
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (A, ) = split_input()
        val (B, ) = split_input()
        val (n, m) = A.length to B.length
        val dp = List(n+1){ MutableList(m+1){false} }
        val preA = getPre(A)
        val preB = getPre(B)

        for(i in 0..n) {
            val a = if(i==0) '*' else A[i-1]
            for (j in 0..m) {
                val b= if(j==0) '*' else B[j-1]
                dp[i][j] = when {
                    i == 0 && j == 0 -> {
                        true
                    }
                    i == 0 -> {
                        dp[i][j - 1] && b == '*'
                    }
                    j == 0 -> {
                        dp[i - 1][j] && a == '*'
                    }
                    a == '*' -> {
                        var fnd = false
                        var k = j
                        for(cnt in 0..4){
                            fnd = fnd || dp[i-1][k]
                            k = preB[k] - 1
                            if(fnd || k < 0) break
                        }
                        fnd
                    }
                    b == '*' -> {
                        var fnd = false
                        var k = i
                        for(cnt in 0..4){
                            fnd = fnd || dp[k][j-1]
                            k = preA[k] - 1
                            if(fnd || k < 0) break
                        }
                        fnd
                    }
                    else -> {
                        a == b && dp[i-1][j-1]
                    }
                }
            }
        }
        val res = if(dp[n][m]) "TRUE" else "FALSE"
        println("Case #${it}: $res")
    }
}

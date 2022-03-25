package google_kickstart.`2017`.C

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val M = 26
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (S, ) = split_input()
        val n = S.length
        val res = if(n % 2 == 1) {
            "AMBIGUOUS"
        }else{
            var res = MutableList(n){' '}
            var pre = 0
            for(i in 0 until n step 2){
                pre = (S[i] - 'A' + M - pre) % M
                res[i+1] = 'A' + pre
            }
            pre = 0
            for(i in n-1 downTo 0 step 2){
                pre = (S[i] - 'A' + M - pre) % M
                res[i-1] = 'A' + pre
            }
            res.joinToString("")
        }
        println("Case #${it}: $res")
    }
}

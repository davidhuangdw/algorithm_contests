package google_kickstart.`2019`.B

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, Q) = input_ints()
        val A = readLine()!!.trim()
        val count = List(N+1){MutableList(26){0} }
        for((i, v) in A.withIndex()){
            for(x in 0 until 26) count[i+1][x] = count[i][x]
            count[i+1][v-'A'] += 1
        }

        fun valid(l: Int, r: Int): Boolean {
            var odd = 0
            for(x in 0 until 26){
                if((count[r][x] - count[l-1][x]) % 2 == 1) {
                    odd += 1
                    if(odd > 1) return false
                }
            }
            return true
        }
        var res = 0
        repeat(Q){
            val (l, r) = input_ints()
            if(valid(l, r))
                res += 1
        }
        println("Case #${it}: $res")
    }
}

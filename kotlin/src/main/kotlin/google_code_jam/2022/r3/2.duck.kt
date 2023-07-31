package google_code_jam.`2022`.r3
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val (N, C) = input_ints()
        val rng = (1..C).map{input_ints()}
        val P = input_ints()
        fun set1(): Long{
            var cnt = 0L
            for(i in 0 until N){
                var miss = hashSetOf<Int>()
                var col = MutableList(N+1){0}
                for(d in 0 until N-1){
                    val j = (i + d) % N
                    val c = P[j]
                    col[c] ++
                    val (a, b) = rng[c-1]
                    if(col[c] !in a..b){
                        if(col[c] > b) break
                        miss.add(c)
                    }else{
                        miss.remove(c)
                    }
                    if(d >= 1 && miss.isEmpty()){
                        cnt ++
                    }
                }
            }
            return cnt
        }
        val r = set1()
        println("Case #${it}: $r")
    }
}

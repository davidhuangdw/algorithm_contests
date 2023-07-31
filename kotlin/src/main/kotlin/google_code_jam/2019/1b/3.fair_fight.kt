package google_code_jam.`2019`.`1b`
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, K) = input_ints()
        val C = input_ints()
        val D = input_ints()
        fun set1(): Int{
            var cnt = 0
            for(l in 0 until N)
                for(r in l until N){
                    if((l..r).maxOf{C[it]} - (l..r).maxOf{D[it]} in -K..K)
                        cnt++
                }
            return cnt
        }
        val res = set1()
        // todo: set2
        println("Case #${it}: $res")
    }
}

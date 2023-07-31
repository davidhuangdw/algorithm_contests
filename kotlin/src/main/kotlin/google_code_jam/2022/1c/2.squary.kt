package google_code_jam.`2022`.`1c`
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val IMP = "IMPOSSIBLE"
    val (T, ) = input_ints()
    (1..T).forEach {
        val (N, K) = input_ints()
        val E = input_longs()
        var pa = 0L
        for(i in 0 until N)
            for(j in i+1 until N)
                pa += E[i] * E[j]
        val sum = E.sum()
        fun set1(): String {
            if(sum == 0L){
                if(pa == 0L)
                    return "1"
                else
                    return IMP
            }else if(pa % sum == 0L){
                return (-pa/sum).toString()
            }else return IMP
        }

        fun set2(): String {
            if(K == 1)
                return set1()
            val a = 1 - sum
            val b = -pa - a*sum
            return "$a $b"
        }
        val r = set2()
        println("Case #${it}: $r")
    }
}

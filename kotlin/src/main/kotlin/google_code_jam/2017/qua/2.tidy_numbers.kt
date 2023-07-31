package google_code_jam.`2017`.qua
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val N = split_input()[0]
        val V = N.map{it - '0'}.toMutableList()
        val n = V.size
        for(i in n-2 downTo 0){
            if(V[i] > V[i+1]){
                V[i] --
                for(j in i+1 until n)
                    V[j] = 9
            }
        }
        var res = 0L
        for(v in V)
            res = res * 10 + v
        println("Case #${it}: $res")
    }
}

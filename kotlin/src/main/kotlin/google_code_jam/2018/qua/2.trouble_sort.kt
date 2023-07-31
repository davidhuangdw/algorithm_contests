package google_code_jam.`2018`.qua
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val (N, ) = input_ints()
        val V = input_ints()
        val part = List(2){mutableListOf<Int>()}
        for((i, v) in V.withIndex()){
            part[i%2].add(v)
        }
        part[0].sort()
        part[1].sort()
        val cur = (0 until N).map{i ->
            part[i%2][i/2]
        }
        fun find(): Int {
            for(i in 0 until N-1){
                if(cur[i] > cur[i+1])
                    return i
            }
            return -1
        }
        val k = find()
        val res = if(k >= 0) k.toString() else "OK"
        println("Case #${it}: $res")
    }
}

package google_code_jam.`2021`.r2
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map{it.toInt()}
    fun input_longs() = split_input().map{it.toLong()}

    val (T, N) = input_ints()
    (1..T).forEach {
        for(i in 1 until N){
            println("M $i $N")
            val j = readLine()!!.toInt()
            if(i != j) {
                println("S $i $j")
                readLine()
            }
        }
        println("D")
        val ok = readLine()!!.toInt()
        assert(ok == 1)
//        println("Case #${it}: ")
    }
}

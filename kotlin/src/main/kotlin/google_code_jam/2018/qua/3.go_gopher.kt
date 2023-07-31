package google_code_jam.`2018`.qua
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val (A, ) = input_ints()
        val K = (A+8)/9 // (k*3 + 1~3, 1~3）
        fun run(){
            for(k in 0 until K){
                val (ci, cj) = k*3 + 2 to 2
                val done = hashSetOf<Pair<Int, Int>>()
                while(done.size < 9){
                    println("$ci $cj")
                    val (i, j) = input_ints()
                    if(i+j == 0)
                        return
                    if(i+j < 0)
                        throw RuntimeException("failed")
                    done.add(i to j)
                }
            }
        }
        run()
//        println("Case #${it}: ")
    }
}

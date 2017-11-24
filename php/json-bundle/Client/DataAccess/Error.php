<?php

namespace DataAccess;

class Error
{

    private $data = [
        'code'    => -1,
        'message' => 'Unknown error',
    ];

    /**
     * @return integer
     */
    public function getCode()
    {
        return $this->data['code'];
    }

    /**
     * @return string
     */
    public function getMessage()
    {
        return $this->data['message'];
    }

    /**
     * @param array $data
     */
    public function setData(array $data)
    {
        $this->setCode($data['code']);
        $this->setMessage($data['message']);
    }

    /**
     * @param int $code
     */
    public function setCode($code)
    {
        $this->data['code'] = (int)$code;
    }

    /**
     * @param string $message
     */
    public function setMessage($message)
    {
        $this->data['message'] = (string)$message;
    }

}
